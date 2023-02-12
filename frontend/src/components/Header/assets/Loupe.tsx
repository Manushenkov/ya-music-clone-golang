export default function Loupe({ className }: { className?: string }) {
	return (
		<svg
			className={className}
			xmlns='http://www.w3.org/2000/svg'
			width='24'
			height='24'
		>
			<path
				d='M19.5 19.2l-3.6-3.7v-.1c.9-1.2 1.5-2.7 1.5-4.3 0-3.9-3.1-7-7-7s-7 3.1-7 7 3.1 7 7 7c1.5 0 2.9-.5 4.1-1.3l.1.1 3.6 3.7.7.7 1.4-1.4-.8-.7zM5.4 11.1c0-2.8 2.2-5 5-5s5 2.2 5 5-2.2 5-5 5-5-2.3-5-5z'
				fill='#fff'
			/>
		</svg>
	)
}
