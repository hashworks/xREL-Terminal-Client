package types

type ReleaseSearchResult struct {
	Total        int          `json:"total"`
	SceneResults []Release    `json:"results"`
	P2PResults   []P2PRelease `json:"p2p_results"`
}

type ExtInfoSearchResult struct {
	Total   int            `json:"total"`
	Results []ShortExtInfo `json:"results"`
}
