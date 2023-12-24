'use client'

import { formatDate } from '@/actions/index'
import Link from 'next/link'
import { BiSolidDislike, BiSolidLike } from 'react-icons/bi'

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
		<div className='flex flex-col w-full max-w-[600px] bg-blue-200 rounded-xl p-2'>
			<div className='mb-4'>
				<Link href={`/posts/${postId}`}>{text}</Link>
			</div>
			<div className='flex flex-row mb-1'>
				<Link className='ml-0 mr-auto' href={`/user/${userId}`}>
					@{username}
				</Link>
				<span className='ml-auto mr-0'>{formatDate(date)}</span>
			</div>
			<div className='flex flex-row gap-6'>
				<div className='flex flex-row items-center gap-1'>
					<div className={isLiked == 'true' ? 'text-green-600' : ''}>
						<BiSolidLike />
					</div>
					{likesCount}
				</div>
				<div className='flex flex-row items-center gap-1'>
					<div className={isDisliked == 'true' ? 'text-red-600' : ''}>
						<BiSolidDislike />
					</div>
					{dislikesCount}
				</div>
			</div>
		</div>
	)
}
