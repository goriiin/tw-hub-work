import Link from 'next/link'

export default function NavBar() {
	const session = null

	return (
		<nav className='flex flex-col items-center text-xl p-4 mt-8 rounded-[0.3rem] fixed border border-gray-200/10 border-l-0'>
			<div className='flex flex-col gap-6'>
				<div className='flex flex-col gap-y-2'>
					<div className='px-6 border-l border-l-gray-200/10 transition-all duration-300 hover:text-primary hover:translate-x-1 hover:border-l-gray-200/30'>
						<Link href='/'>Home</Link>
					</div>
					<div className='px-6 border-l border-l-gray-200/10 transition-all duration-300 hover:text-primary hover:translate-x-1 hover:border-l-gray-200/30'>
						<Link href='/feed'>Feed</Link>
					</div>
				</div>
				<div className='px-6 border-l border-l-gray-200/10 transition-all duration-300 hover:text-primary hover:translate-x-1 hover:border-l-gray-200/30'>
					<div>{session ? 'Profile' : <Link href='/auth'>Sign up</Link>}</div>
				</div>
			</div>
		</nav>
	)
}
