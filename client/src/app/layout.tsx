import Header from '@/components/Header'
import NavBar from '@/components/NavBar'
import type { Metadata } from 'next'
import { Suspense } from 'react'
import './globals.css'

export const metadata: Metadata = {
	title: 'Twit Hub',
	description: 'Twit Hub',
}

export default function RootLayout({
	children,
}: {
	children: React.ReactNode
}) {
	return (
		<html lang='en'>
			<body className='bg-stone-950 text-zinc-300'>
				<div>
					<Header />
					<div className='flex flex-col'>
						<div className='shrink-0'>
							<NavBar />
						</div>
						<div className='grow'>
							<Suspense fallback={<div>Loading...</div>}>{children}</Suspense>
						</div>
					</div>
				</div>
			</body>
		</html>
	)
}
