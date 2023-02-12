export default function Plus({ className }: { className?: string }) {
	return (
		<svg
			className={className}
			xmlns='http://www.w3.org/2000/svg'
			fillRule='evenodd'
			clipRule='evenodd'
			strokeLinejoin='round'
			strokeMiterlimit='1.414'
		>
			<path
				d='M11 11H0v2h11v11h2V13h11v-2H13V0h-2v11z'
				fill='#fff'
				fillRule='nonzero'
			/>
		</svg>
	)
}
