import data from '@/data/data.json'
import Post from './Post'

export default function FeedPage() {
	return (
		<div className='flex flex-col justify-center text-center items-center max-w-[600px] gap-y-2'>
			{data.posts.map((post) => (
				<Post
					key={post.id}
					id={post.id}
					title={post.title}
					body={post.body}
					userId={post.userId}
					tags={post.tags}
					reactions={post.reactions}
				/>
			))}
		</div>
	)
}
