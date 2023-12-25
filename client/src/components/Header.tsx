import Link from 'next/link'

export default function Header() {
	const session = null

	return (
		<div className='flex flex-row px-4 py-3 text-l bg-zinc-800'>
			<div className='ml-0 mr-auto'>Menu</div>
			<div className='flex flex-row gap-10 ml-auto mr-0'>
				<div>Seach Bar</div>
				<div>{session ? 'Profile' : <Link href='/auth'>Sign up</Link>}</div>
			</div>
		</div>
	)
}
