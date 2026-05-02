package main

import (
  "os"
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

func main() {
	// Skeleton of a WiiU/3DS Friends server running on PRUDPv0 with a single endpoint

	authServer := nex.NewPRUDPServer() // The main PRUDP server
	endpoint := nex.NewPRUDPEndPoint(1) // A PRUDP endpoint for PRUDP connections to connect to. Bound to StreamID 1
	endpoint.ServerAccount = nex.NewAccount(types.NewPID(1), "Quazal Authentication", "password"))
	endpoint.AccountDetailsByPID = accountDetailsByPID
	endpoint.AccountDetailsByUsername = accountDetailsByUsername

	// Setup event handlers for the endpoint
	endpoint.OnData(func(packet nex.PacketInterface) {
		if packet, ok := packet.(nex.PRUDPPacketInterface); ok {
			request := packet.RMCMessage()

			fmt.Println("[AUTH]", request.ProtocolID, request.MethodID)

			if request.ProtocolID == 0xA { // TicketGrantingProtocol
				if request.MethodID == 0x1 { // TicketGrantingProtocol::Login
					handleLogin(packet)
				}

				if request.MethodID == 0x3 { // TicketGrantingProtocol::RequestTicket
					handleRequestTicket(packet)
				}
			}
		}
	})

	// Bind the endpoint to the server and configure it's settings
	authServer.BindPRUDPEndPoint(endpoint)
	authServer.SetFragmentSize(962)
	authServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(1, 1, 0))
	authServer.SessionKeyLength = 16
	authServer.AccessKey = "ridfebb9"
	authServer.Listen(os.Getenv("PN_YKWB_SECURE_SERVER_PORT"))
}
