/* eslint-disable jsx-a11y/alt-text */
import './UserPlaylists.scss'
import Plus from './assets/Plus'
import playlistCoverLike from './assets/playlistCoverLike.png'
import playlistCover from './assets/playlistCover.png'
import { Dispatch, SetStateAction, useEffect, useState } from 'react'
import classNames from 'classnames'

interface UserPlaylistsProps {
	selectedPlaylist: PlaylistData
	setSelectedPlaylist: Dispatch<SetStateAction<PlaylistData>>
	playlists: PlaylistData[]
	setPlaylists: Dispatch<SetStateAction<PlaylistData[]>>
}

export interface PlaylistData {
	name: string
	id: string
}

export default function UserPlaylists({
	selectedPlaylist,
	setSelectedPlaylist,
	playlists,
	setPlaylists,
}: UserPlaylistsProps) {
	const [isShowPlaylistCreation, setIsShowPlaylistCreation] = useState(false)
	const [playlistCreationName, setPlaylistCreationName] = useState('')

	useEffect(() => {
		fetch('http://localhost:8080/user/playlists')
			.then((r) => r.json())
			.then(setPlaylists)
	}, [setPlaylists])

	return (
		<div className='user-playlists'>
			<h2 className='user-playlists__title'>Плейлисты</h2>
			<div className='user-playlists__cards'>
				<div
					onClick={() =>
						setSelectedPlaylist({ id: '0', name: 'Мне нравится' })
					}
					className={classNames('user-playlists__card', {
						'user-playlists__card-selected':
							'0' === selectedPlaylist.id,
					})}
				>
					<img src={playlistCoverLike}></img>
					<div>Мне нравится</div>
				</div>
				<div
					onClick={() => {
						setIsShowPlaylistCreation(true)
					}}
					className='user-playlists__plus-card'
				>
					{isShowPlaylistCreation ? (
						<div>
							<button
								onClick={() =>
									fetch(
										'http://localhost:8080/playlist/create/' +
											playlistCreationName,
										{ method: 'POST' }
									)
								}
							>
								Create
							</button>
							<button
								onClick={(e) => {
									e.stopPropagation()
									setIsShowPlaylistCreation(false)
								}}
							>
								Cancel
							</button>
							<input
								type='text'
								value={playlistCreationName}
								onChange={(e) =>
									setPlaylistCreationName(e.target.value)
								}
							/>
						</div>
					) : (
						<Plus className='user-playlists__plus' />
					)}
				</div>
				{playlists.map((playlist) => (
					<div
						key={playlist.id}
						onClick={() => setSelectedPlaylist(playlist)}
						className={classNames('user-playlists__card', {
							'user-playlists__card-selected':
								playlist.id === selectedPlaylist.id,
						})}
					>
						<img src={playlistCover}></img>
						{playlist.name}
					</div>
				))}
			</div>
		</div>
	)
}
