'use client'

import { formatDate } from '@/actions/index'
import Link from 'next/link'

type Props = {
	postId: number
	userId: number
	username: string
	text: string
	date: string
	isLiked: string
	likesCount: string
	isDisliked: string
	dislikesCount: string
}

export default function Post({
	postId,
	userId,
	username,
	text,
	date,
	isLiked,
	likesCount,
	isDisliked,
	dislikesCount,
}: Props) {
	return (
		<div className='flex flex-col gap-4 w-full max-w-[600px] bg-blue-200 rounded-xl p-2'>
			<Link href={`/posts/${postId}`}>{text}</Link>
			<div className='flex flex-row gap-8'>
				<Link className='ml-0 mr-auto' href={`/user/${userId}`}>
					@{username}
				</Link>
				<span className='ml-auto mr-0'>{formatDate(date)}</span>
			</div>
		</div>
	)
}
