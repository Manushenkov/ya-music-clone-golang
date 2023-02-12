import Loupe from '../assets/Loupe'
import './Controls.scss'

export default function Controls() {
	return (
		<div className='controls'>
			<div className='controls__navigation'>
				<div>Главное</div>
				<div>Подкасты и книги</div>
				<div>Детям</div>
				<div>Потоки</div>
				<div>Коллекция</div>
			</div>
			<div className='controls__search'>
				<Loupe className='header__logo' />
			</div>
		</div>
	)
}
