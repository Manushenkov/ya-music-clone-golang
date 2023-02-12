export default function Dots({ className }: { className?: string }) {
	return (
		<svg
			className={className}
			xmlns='http://www.w3.org/2000/svg'
			width='24'
			height='24'
		>
			<path
				d='M6 13.5c-.8 0-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5 1.5.7 1.5 1.5-.7 1.5-1.5 1.5zm6 0c-.8 0-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5 1.5.7 1.5 1.5-.7 1.5-1.5 1.5zm6 0c-.8 0-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5 1.5.7 1.5 1.5-.7 1.5-1.5 1.5z'
				fill='#fff'
			/>
		</svg>
	)
}
