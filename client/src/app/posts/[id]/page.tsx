type Props = {
	params: {
		postId: number
	}
}

export default function User({ params }: Props) {
	return (
		<div>
			<h1>Post Page</h1>
			<h2>{params.postId}</h2>
		</div>
	)
}
