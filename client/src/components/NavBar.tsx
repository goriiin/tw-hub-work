import Link from 'next/link'

export default function NavBar() {
	const session = null

	return (
		<nav className='flex flex-col items-center text-xl gap-8 px-16 py-8 mt-8 rounded-[0.3rem] fixed bg-zinc-800'>
			<div className='flex flex-col gap-y-2'>
				<Link href='/'>Home</Link>
				<Link href='/feed'>Feed</Link>
			</div>
			<div>
				<div>{session ? 'Profile' : <Link href='/auth'>Sign up</Link>}</div>
			</div>
		</nav>
	)
}
