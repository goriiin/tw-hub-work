import data from '@/data/data.json'
import Post from './Post'

export default function FeedPage() {
	// axios.get('http://localhost:8000/errrm/what/da/flip').then((response) => {
	// 	console.log(response.data)
	// })

	return (
		<div className='flex flex-col justify-center items-center gap-y-2 mt-8'>
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
