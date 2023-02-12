import { Dispatch, SetStateAction, useEffect, useRef, useState } from 'react'
import './Playlist.scss'
import Cross from './assets/Cross'
import Dots from './assets/Dots'
import Play from './assets/Play'
import { PlaylistData } from '../UserPlaylists/UserPlaylists'

interface PlaylistProps {
	selectedPlaylist: PlaylistData
	setSelectedPlaylist: any
	setPlaylists: Dispatch<SetStateAction<PlaylistData[]>>
}

interface Track {
	name: string
	id: string
}

export default function Playlist({
	selectedPlaylist,
	setSelectedPlaylist,
	setPlaylists,
}: PlaylistProps) {
	const trackInput = useRef<HTMLFormElement>(null)

	const [tracks, setTracks] = useState<Track[]>([])

	useEffect(() => {
		if (selectedPlaylist.id === '0') {
			setTracks([])
			return
		}
		fetch(`http://localhost:8080/playlist/${selectedPlaylist.id}/tracks`)
			.then((r) => r.json())
			.then((data) => setTracks(data.tracks))
	}, [selectedPlaylist])

	const playTrack = async (id: string) => {
		const ctx = new AudioContext()
		// const gainNode = ctx.createGain()
		const track = await fetch('http://localhost:8080/track/' + id)
			.then((data) => data.arrayBuffer())
			.then((arrayBuffer) => ctx.decodeAudioData(arrayBuffer))

		const playAudio = ctx.createBufferSource()
		playAudio.buffer = track
		// playAudio.connect(ctx.destination)
		playAudio.start(ctx.currentTime)

		const gainNode = ctx.createGain()
		gainNode.connect(ctx.destination)
		playAudio.connect(gainNode)
		gainNode.gain.value = 0.05
	}

	const deleteTrack = (id: string) => {
		fetch(
			'http://localhost:8080/playlist/' +
				selectedPlaylist.id +
				'/tracks/delete/' +
				id,
			{ method: 'DELETE' }
		)
		setTracks((prev) => prev.filter((track) => track.id !== id))
	}

	const deletePlaylist = (id: string) => {
		fetch('http://localhost:8080/playlist/delete/' + id, {
			method: 'DELETE',
		}).then(() => {
			setSelectedPlaylist({ id: '0', name: '' })
			setPlaylists((prev) =>
				prev.filter((playlist) => playlist.id !== id)
			)
		})
	}

	return (
		<div className='playlist'>
			<Cross className='playlist__cross' />
			<div className='playlist__playlist'>Плейлист</div>
			<div className='playlist__name'>{selectedPlaylist.name}</div>
			<div className='playlist__author'>Автор: artem.manushenkov</div>
			<div className='playlist__controls'>
				<button className='playlist__play-button'>
					<Play className='playlist__play-icon' />
				</button>
				<button className='playlist__dots-button'>
					<Dots className='playlist__dots-icon' />
				</button>
				<button onClick={() => deletePlaylist(selectedPlaylist.id)}>
					Удалить текущий плейлист
				</button>
			</div>
			<form
				ref={trackInput}
				id='formElem'
				method='post'
				encType='multipart/form-data'
				onSubmit={(e) => {
					e.preventDefault()
					if (!trackInput.current) return

					fetch(
						'http://localhost:8080/playlist/' +
							selectedPlaylist.id +
							'/tracks/upload',
						{
							method: 'POST',

							body: new FormData(trackInput.current),
						}
					)
				}}
			>
				<input
					type='file'
					defaultValue={undefined}
					name='file'
					accept='audio/*'
				/>
				<input type='submit' />
			</form>
			{tracks.map((track, index) => (
				<div key={track.id} className='track'>
					<span>{index} </span>
					<span>{track.name} </span>
					<span>{track.id}</span>
					<button onClick={() => playTrack(track.id)}>play</button>
					<button onClick={() => deleteTrack(track.id)}>
						delete
					</button>
				</div>
			))}
		</div>
	)
}
