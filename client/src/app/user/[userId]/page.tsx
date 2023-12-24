type Props = {
	params: {
		userId: number
	}
}

export default function User({ params }: Props) {
	return (
		<div>
			<h1>User Page</h1>
			<h2>{params.userId}</h2>
		</div>
	)
}
