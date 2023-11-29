import Link from 'next/link'

export default function Home() {
	return (
		<div>
			<h1>Do not have a profile? Sing up:</h1>
			<Link href='/auth'>Sign up</Link>
		</div>
	)
}
