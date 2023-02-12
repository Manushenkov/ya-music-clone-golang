import Controls from './Controls/Controls'
import './Header.scss'
import Logo from './assets/Logo'

export default function Header() {
	return (
		<header className='header'>
			<Logo className='header__logo' />
			<Controls />
		</header>
	)
}
