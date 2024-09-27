package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"
)

func NewMPDServer(address string, requests chan ApiRequest) (*MPDServer, error) {
	s := &MPDServer{
		startTime: time.Now(),
		requests:  requests,
	}

	var err error
	s.listener, err = net.Listen("tcp", address)
	if err != nil {
		return s, fmt.Errorf("failed starting api listener: %w", err)
	}

	log.Infof("MPD server listening on %s", s.listener.Addr())

	go s.serve()
	return s, nil
}

type MPDServer struct {
	startTime   time.Time
	listener    net.Listener
	requests    chan ApiRequest
	jobid       atomic.Uint64
	clients     []MPDConn
	clientsLock sync.Mutex
}

type MPDConn struct {
	Events atomic.Uint64
}

const (
	mpdEventDatabase = 1 << iota
	mpdEventUpdate
	mpdEventPlaylist
	mpdEventPlayer
	mpdEventMixer
	mpdEventOptions
)

func (s *MPDServer) serve() {
	for {
		// Accept a client connection.
		conn, err := s.listener.Accept()
		if err != nil {
			// Not sure what to do here otherwise.
			log.WithError(err).Errorln("failed to accept new MPD connection, shutting down MPD server")
			s.listener.Close()
			break
		}

		mpdconn := &MPDConn{}

		go func() {
			log.Infof("accepting MPD client connection %s", conn.RemoteAddr())
			err := s.serveConn(conn, mpdconn)
			if err != nil {
				log.WithError(err).Errorf("MPD connection with %s failed, closing connection", conn.RemoteAddr())
			}
		}()
	}
}

func (s *MPDServer) serveConn(conn net.Conn, mpdconn *MPDConn) error {
	_, err := fmt.Fprintf(conn, "OK MPD 0.24\n")
	if err != nil {
		return fmt.Errorf("failed to send initial version string: %w", err)
	}

	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("failed to read command from client: %w", err)
		}
		line = strings.TrimRight(line, "\r\n")
		fmt.Println("line:", line)

		// Read command name.
		command, args := s.parseCommand(line)

		switch command {
		// Querying MPDs status
		case "currentsong":
			fmt.Fprintf(conn, "OK\n")
		case "idle":
			// TODO: implement
		case "status":
			req := ApiRequest{
				Type: ApiRequestTypeStatus,
				resp: make(chan apiResponse, 1),
			}
			s.requests <- req
			resp := <-req.resp
			if resp.err != nil {
				fmt.Fprintf(conn, "ACK [5@0] %s error: %s\n", command, resp.err.Error())
				continue
			}
			status := resp.data.(*ApiResponseStatus)

			repeat := 0
			if status.RepeatContext {
				repeat = 1
			}
			random := 0
			if status.ShuffleContext {
				random = 1
			}
			state := "play"
			if status.Stopped {
				state = "stop"
			} else if status.Paused {
				state = "pause"
			}
			fmt.Fprintf(conn, "repeat: %d\n", repeat)
			fmt.Fprintf(conn, "random: %d\n", random)
			fmt.Fprintf(conn, "single: 0\n")
			fmt.Fprintf(conn, "consume: 0\n")
			fmt.Fprintf(conn, "partition: default\n")
			fmt.Fprintf(conn, "playlist: 1\n")
			fmt.Fprintf(conn, "playlistlength: 0\n")
			fmt.Fprintf(conn, "mixrampdb: 0\n")
			fmt.Fprintf(conn, "state: %s\n", state)
			fmt.Fprintf(conn, "OK\n")

		case "stats":
			fmt.Fprintf(conn, "artists: 0\n")
			fmt.Fprintf(conn, "albums: 0\n")
			fmt.Fprintf(conn, "songs: 0\n")
			fmt.Fprintf(conn, "uptime: %d\n", int64(time.Since(s.startTime).Seconds()))
			fmt.Fprintf(conn, "db_playtime: 0\n")
			fmt.Fprintf(conn, "db_uptime: 0\n")
			fmt.Fprintf(conn, "playtime: 0\n")
			fmt.Fprintf(conn, "OK\n")

		// Queue
		case "playlistinfo":
			// TODO: nothing in the playlist yet
			fmt.Fprintf(conn, "OK\n")

		// Stored playlists
		case "listplaylistinfo":
			fmt.Fprintf(conn, "ACK [50@0] {listplaylistinfo} No such playlist")

		// Stored playlists
		case "listplaylists":
			fmt.Fprintf(conn, "OK\n")

		// Music database
		case "list":
			// Nothing to list here.
			fmt.Fprintf(conn, "OK\n")
		case "lsinfo":
			// Every directory is empty.
			//fmt.Fprintf(conn, "ACK [2@0] {lsinfo} Unsupported URI scheme\n")
			//fmt.Fprintf(conn, "ACK [50@0] {lsinfo} No such directory\n")
			fmt.Fprintf(conn, "OK\n")
		case "update":
			fmt.Fprintf(conn, "updating_db: %d\n", s.jobid.Add(1))
			fmt.Fprintf(conn, "OK\n")

			// The database is immediately updated, with no changes.
			s.emitEvents(mpdEventUpdate)

		// Connection settings
		case "close":
			conn.Close()
			return nil
		case "tagtypes":
			// No tag types available.
			fmt.Fprintf(conn, "OK\n")

		// Partition commands
		case "listpartitions":
			fmt.Fprintf(conn, "partition: default\n")
			fmt.Fprintf(conn, "OK\n")
		case "partition":
			if len(args) != 1 {
				fmt.Fprintf(conn, "ACK [2@0] {status} wrong number of arguments\n")
				continue
			}
			if args[0] != "default" {
				fmt.Fprintf(conn, "ACK [50@0] {partition} partition does not exist\n")
				continue
			}
			fmt.Fprintf(conn, "OK\n")

		// Audio output devices
		case "outputs":
			fmt.Fprintf(conn, "outputid: 0\n")
			fmt.Fprintf(conn, "outputname: Default output\n") // TODO: show configured output
			fmt.Fprintf(conn, "outputenabled: 0\n")
			fmt.Fprintf(conn, "OK\n")

		// Reflection
		case "config":
			// Don't show configuration.
			fmt.Fprintf(conn, "OK\n")
		case "commands":
			// List available commands.
			fmt.Fprintf(conn, "OK\n")
		case "urlhandlers":
			// No URL handlers available.
			fmt.Fprintf(conn, "OK\n")
		case "decoders":
			// No decoders available.
			fmt.Fprintf(conn, "OK\n")

		default:
			fmt.Printf("unknown command: %s %v\n", command, args)
			fmt.Fprintf(conn, "ACK [5@0] %s Unknown command\n", command)
		}
	}
}

func (s *MPDServer) parseCommand(line string) (string, []string) {
	// Parse command itself.
	index := strings.IndexByte(line, ' ')
	if index < 0 {
		return line, nil // no arguments
	}
	command := line[:index]

	// Parse arguments
	var args []string
	escape := false
	quoted := false
	arg := ""
	for _, c := range line[index:] {
		if c == ' ' {
			if quoted {
				arg += string(c)
			} else if arg != "" {
				// end of arg
				args = append(args, arg)
				arg = ""
			}
		} else if !escape && c == '"' {
			if quoted {
				// end of arg
				args = append(args, arg)
				arg = ""
			} else {
				// beginning of arg
				quoted = true
			}
		} else if c == '\\' {
			// TODO: deal with escape chars
		} else {
			arg += string(c)
		}
		escape = c == '\\'
	}
	if arg != "" {
		args = append(args, arg)
	}

	return command, args
}

func (s *MPDServer) Emit(ev *ApiEvent) {
	var events uint64
	switch ev.Type {
	case ApiEventTypeWillPlay:
	case ApiEventTypePlaying, ApiEventTypeNotPlaying, ApiEventTypeStopped, ApiEventTypePaused:
		events |= mpdEventPlayer
	default:
		fmt.Printf("unknown event: %s\n", ev.Type)
	}
	if events != 0 {
		s.emitEvents(events)
	}
}

func (s *MPDServer) emitEvents(events uint64) {
	// TODO
}
