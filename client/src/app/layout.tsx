import NavBar from '@/components/NavBar'
import type { Metadata } from 'next'
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
			<body>
				<NavBar />
				{children}
			</body>
		</html>
	)
}
