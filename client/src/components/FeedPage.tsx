import data from '@/data/data.json'
import Post from './Post'

export default function FeedPage() {
	return (
		<div className='flex flex-col justify-center items-center gap-y-2'>
			{data.posts.map((post) => (
				<Post
					key={post.postId}
					postId={post.postId}
					userId={post.userId}
					username={post.username}
					text={post.text}
					date={post.date}
					isLiked={post.isLiked}
					likesCount={post.likesCount}
					isDisliked={post.isDisliked}
					dislikesCount={post.dislikesCount}
				/>
			))}
		</div>
	)
}
