package types

type FavList struct {
	Id				int		`json:"id"`
	Name			string	`json:"name"`
	IsPublic		bool	`json:"public"`
	DoNotify		bool	`json:"notify"`
	DoAutoRead		bool	`json:"auto_read"`
	IncludesP2P		bool	`json:"include_p2p"`
	Description		string	`json:"description"`
	PasswordHash	string	`json:"passwort_hash"`		// $password_hash = sha1($list->id . "\r\n" . $list->password);
	EntryCount		int		`json:"entry_count"`
	UnreadReleases	int		`json:"unread_releases"`
}

type ShortFavList struct {
	Id				string	`json:"id"`					// TODO: Why is this a string?
	Name			string	`json:"string"`
	Entries			string	`json:"entries"` 			// TODO: Why is this a string?
	UnreadReleases	string	`json:"unread_releases"`	// TODO: Why is this a string?
}

type FavListEntryModificationResult struct {
	Success	int				`json:"success"`			// TODO: Why is this an int?
	FavList	ShortFavList	`json:"fav_list"`
	ExtInfo	ShortExtInfo	`json:"ext_info"`
}