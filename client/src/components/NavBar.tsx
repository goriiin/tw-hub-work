import Link from 'next/link'

export default function NavBar() {
	return (
		<nav className='flex flex-row place-content-between text-l py-4'>
			<div className='ml-4'>Logo</div>
			<div className='flex gap-2'>
				<Link href='/'>Home</Link>
				<Link href='/feed'>Feed</Link>
			</div>
			<div className='mr-4'>
				<Link href='/auth'>Sign up</Link>
			</div>
		</nav>
	)
}
