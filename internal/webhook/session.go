package webhook

import (
	launchutils "github.com/s1dorfake/HypurrFun/internal/launch_utils"
	hp "github.com/s1dorfake/HypurrFun/pb"
)

const notAvailable string = "Not Available"

// getDefaultSession returns a default session when none is provided
func getDefaultSession() *hp.HyperliquidWalletDeploySession {
	return &hp.HyperliquidWalletDeploySession{
		FullName:       notAvailable,
		TokenName:      notAvailable,
		TokenSupply:    0,
		StartMarketCap: 0,
	}
}

// getSessionInfo safely extracts session information
func getSessionInfo(launch *launchutils.LaunchExtended) *hp.HyperliquidWalletDeploySession {
	if launch.Session == nil {
		return getDefaultSession()
	}
	return launch.Session
}
