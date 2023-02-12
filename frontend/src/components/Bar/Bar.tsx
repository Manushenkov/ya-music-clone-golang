import './Bar.scss'
import Next from './assets/Next'
import Play from './assets/Play'
import Prev from './assets/Prev'

export default function Header() {
	return (
		<div className='bar'>
			<div className='bar__progress'></div>
			<div className='bar__controls'>
				<div className='bar__btn'>
					<Prev />
				</div>
				<div className='bar__btn'>
					<Play />
				</div>
				<div className='bar__btn'>
					<Next />
				</div>
			</div>
		</div>
	)
}
