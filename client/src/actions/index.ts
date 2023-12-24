export function formatDate(s: string): string {
	const formatted = new Date(s)

	const daysOfWeek = [
		'Sunday',
		'Monday',
		'Tuesday',
		'Wednesday',
		'Thursday',
		'Friday',
		'Saturday',
	]

	const months = [
		'January',
		'February',
		'March',
		'April',
		'May',
		'June',
		'July',
		'August',
		'September',
		'October',
		'November',
		'December',
	]

	const dayOfMonth = formatted.getDate().toString().padStart(2, '0')
	const dayOfWeek = daysOfWeek[formatted.getDay()]
	const month = months[formatted.getMonth()]
	const hours = formatted.getHours().toString().padStart(2, '0')
	const minutes = formatted.getMinutes().toString().padStart(2, '0')

	return `${dayOfWeek}, ${dayOfMonth} ${month} ${hours}:${minutes}`
}
