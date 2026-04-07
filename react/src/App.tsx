import { useState, useEffect } from 'react';
import axios from 'axios';
import "./App.css";

type Event = {
	id: number;
	user_id: number;
	action: string;
	timestamp: string;
	metadata?: Record<string, any>;
};

type Response = { events: Event[] } | { error: string };

function App() {
	const [events, setEvents] = useState<Event[]>([]);
	const [userId, setUserId] = useState<number | string>("");
	const [startTime, setStartTime] = useState<string>("");
	const [endTime, setEndTime] = useState<string>("");

	async function getEvents() {
		const params: Record<string, string | number> = {};
		if (userId) {
			params.user_id = userId;
		}
		if (startTime) {
			params.start_time = new Date(startTime).toISOString();
		} 
		if (endTime) {
			params.end_time = new Date(endTime).toISOString();
		}

		try {
			const response = await axios.get<Response>('http://localhost:8000/event', { params: params });

			if ('error' in response.data) {
				console.error(response.data.error);
			} else {
				setEvents(response.data.events ?? []);
			}
		} catch (error: any) {
			console.error(error.message || error);
		}
	}

	useEffect(() => {
		getEvents();
	}, []);

  	return (
    	<>
			<div className='app'>
				<div className='filters'>
					<input type='number'
						value={userId}
						onChange={(e) => setUserId(e.target.value ? Number(e.target.value) : "")}
					></input>
					<input type='datetime-local'
						value={startTime}
						onChange={(e) => setStartTime(e.target.value)}
					></input>
					<input type='datetime-local'
						value={endTime}
						onChange={(e) => setEndTime(e.target.value)}
					></input>
					<button onClick={getEvents}>Find</button>
				</div>
				<div>
					<table className='events'>
						<thead>
							<tr>
								<th>User Id</th>
								<th>Action</th>
								<th>Metadata</th>
								<th>Timestamp</th>
							</tr>
						</thead>
						<tbody>
							{events.map((event) => (
								<tr key={event.id}>
									<td>{event.user_id}</td>
									<td>{event.action}</td>
									<td>{event.metadata ? JSON.stringify(event.metadata) : ''}</td>
									<td>{new Date(event.timestamp).toLocaleString()}</td>
								</tr>
							))}
						</tbody>
					</table>
				</div>
			</div>
    	</>
  	)
}

export default App
