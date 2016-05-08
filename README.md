# xREL Terminal Client

A terminal client to access the [xREL.to](http://xrel.to) API, written in Go.

## Install

Simply download the latest release for your arch from [here](https://github.com/hashworks/xRELTerminalClient/releases/latest) and execute it.
On your first start it will create a config file, you can avoid that by setting `--configFile=/dev/null`.
For global access place the executable in your `$PATH`.

## Usage

See `--help`:
```
Global flags:

--configFile /path/to/config.json
	Sets the path to the config file to use.
	Don't want a config file? Set it to /dev/null.

--p2p
	Shows P2P instead of scene results.
	Basically it can be used with every function
	that shows releases.

--perPage 5
	Set how many entries to show per page.
	Basically it can be used with every function
	that displays pagination in any way.

--page 1
	Set the page to show. Can be used along with --perPage.

Function flags:

--version
	Shows the version and a few informations.

--authenticate
	Authenticates you with xREL.to using oAuth2.

--rateLimit
	Shows your current rate limit.

--searchRelease "Portal 2 Linux"
	Search for a release. Optional parameters:
	--limit 5
		Limit output from 5 to 25 entries.
		Uses value of --perPage by default.
	
--release Portal.2.Linux-ACTiVATED
	Show information about a release.
	Optional parameters, all of them require authentication:
	--addComment "[...]"
		Add a comment to a release.
	--rateVideo 9
		Rate the video of a release from 1-10. Requires --rateAudio.
	--rateAudio 8
		Rate the audio of a release from 1-10. Requires --rateVideo.

--comments Portal.2.Linux-ACTiVATED
	List comments of a release.

--getNFOImage Portal.2.Linux-ACTiVATED
	Saves an image of the NFO of the specified release
	in the current directory.

--addProof=filepathTo/Proof/image.png Game.of.Thrones.S05E10.Die.Gnade.der.Mutter.German.DL.1080p.BluRay.x264-RSG Game.of.Thrones.S05E10.Die.Gnade.der.Mutter.German.DL.720p.BluRay.x264-RSG
	Adds a proof image to the specified releases. Requires authentication.

--searchMedia "The Big Bang Theory"
	Search for media. Optional parameters:
	--mediaType tv
		Limit results by movie, tv, game, console, software or xxx.
	--limit 5
		See --searchRelease.
	--addToFavorites
		Add selected media to a favorites list you select.
		Requires authentication. Optional parameters:
		--listName Games
			Select a specific favorites list.
	--info
		Show information about the selected media.
		Only usefull if used with the following parameters.
	--rate 8
		Rate selected media from 1 to 10.
	--releases
		Show latest releases of the selected media.
	--images
		Show images of the selected media.
	--videos
		Show videos of the selected media.

--showUnreadFavorites
	Select a user's favorite list and show unread releases.
	Requires authentication. Optional parameters:
	--listName Games
		Display only a specific list.
	--markAsRead
		Marks entries as read.

--removeFavoriteEntry
	Select a user's favorite list and remove an entry.
	Requires authentication. Optional parameters:
	--listName Games
		Select a specific list.

--latest
	Lists latest releases. Optional parameters:
	--filter overview
		Filter ID or 'overview' to use the currently
		logged in user's overview filter.

--browseArchive YYYY-MM
	Browse archive. Optional parameters:
	--filter overview
		See --latest.

--filters
	Shows a list of public, predefined release filters
	to use with --filter.

--browseCategory topmovie
	Browse a category. Optional parameters:
	--mediaType movie
		See --searchMedia.

--categories
	Returns a list of categories to use with --browseCategory.

--upcomingTitles
	Lists upcoming titles. Optional parameters:
	--country us
		Show upcoming titles for a specific country, currently
		possible values are 'de' (default) and 'us'.
	--releases
		See --searchMedia.
```
