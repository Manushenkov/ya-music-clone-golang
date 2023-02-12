import { useState } from 'react'
import './App.scss'
import Bar from './components/Bar/Bar'
import Header from './components/Header/Header'
import Playlist from './components/Playlist/Playlist'
import UserPlaylists, {
	PlaylistData,
} from './components/UserPlaylists/UserPlaylists'

function App() {
	const [selectedPlaylist, setSelectedPlaylist] = useState<PlaylistData>({
		id: '0',
		name: '',
	})
	const [playlists, setPlaylists] = useState<PlaylistData[]>([])

	return (
		<div className='app'>
			<Header />
			<div className='app__content'>
				<div className='app__main'>
					<UserPlaylists
						selectedPlaylist={selectedPlaylist}
						setSelectedPlaylist={setSelectedPlaylist}
						playlists={playlists}
						setPlaylists={setPlaylists}
					/>
				</div>
				<Playlist
					selectedPlaylist={selectedPlaylist}
					setSelectedPlaylist={setSelectedPlaylist}
					setPlaylists={setPlaylists}
				/>
			</div>
			<Bar />
		</div>
	)
}

export default App
