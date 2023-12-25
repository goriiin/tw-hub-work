import Link from 'next/link'

export default function Home() {
	return (
		<div className='flex'>
			<h1>Do not have a profile?</h1>
			<Link href='/auth'>Sign up</Link>
		</div>
	)
}
