'use client'

type Props = {
	id: number
	title: string
	body: string
	userId: number
	tags: string[]
	reactions: number
}

export default function Post({
	id,
	title,
	body,
	userId,
	tags,
	reactions,
}: Props) {
	return (
		<div className='flex flex-col bg-blue-200 rounded-xl text-center p-4'>
			<h2 className='text-xl'>{title}</h2>
			<p>{body}</p>
			<div className='flex flex-row gap-1 justify-center'>
				Tags:
				{tags.map((tag, i) => (
					<span key={i}>{tag}</span>
				))}
			</div>
			<div>Likes: {reactions}</div>
		</div>
	)
}
