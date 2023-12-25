'use client'

import { formatDate } from '@/actions/index'
import Link from 'next/link'
import { BiSolidDislike, BiSolidLike } from 'react-icons/bi'
import { FaCommentAlt } from 'react-icons/fa'

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
		<div className='flex flex-col w-full max-w-[600px] bg-zinc-800 rounded-[0.3rem] p-2'>
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
					<div className={isLiked == 'true' ? 'text-green-700' : ''}>
						<BiSolidLike />
					</div>
					<span>{likesCount}</span>
				</div>
				<div className='flex flex-row items-center gap-1'>
					<div className={isDisliked == 'true' ? 'text-rose-700' : ''}>
						<BiSolidDislike />
					</div>
					<span>{dislikesCount}</span>
				</div>
				<div className='flex flex-row items-center gap-1'>
					<Link href={`/posts/${postId}`}>
						<div className='flex flex-row items-center gap-1'>
							<span>Comments</span>
							<FaCommentAlt />
						</div>
					</Link>
				</div>
			</div>
		</div>
	)
}
