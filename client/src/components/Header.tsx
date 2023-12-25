import Link from 'next/link'

export default function Header() {
	const session = null

	return (
		<div className='flex flex-row -40 px-4 py-3 text-l sticky top-0 backdrop-filter backdrop-blur bg-opacity-80 border-b border-b-gray-200/10 bg-background'>
			<div className='ml-0 mr-auto'>Menu</div>
			<div className='flex flex-row gap-10 ml-auto mr-0'>
				<div>
					<input
						type='text'
						className='border border-gray-200/10 bg-background rounded block w-full p-1
						outline-none focus:border-stone-200/30'
						placeholder='Find...'
					/>
				</div>
				<div>{session ? 'Profile' : <Link href='/auth'>Sign up</Link>}</div>
			</div>
		</div>
	)
}
