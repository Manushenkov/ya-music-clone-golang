export default function Cross({ className }: { className?: string }) {
	return (
		<svg
			className={className}
			xmlns='http://www.w3.org/2000/svg'
			width='24'
			height='24'
		>
			<path
				fill='#fff'
				d='M13.414 12l4.95-4.95-1.414-1.414-4.95 4.95-4.95-4.95L5.636 7.05l4.95 4.95-4.95 4.95 1.414 1.414 4.95-4.95 4.95 4.95 1.414-1.414-4.95-4.95z'
			/>
		</svg>
	)
}
