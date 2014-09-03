<h1>Places</h1>

<table class="placesOverview pure-table">
	<thead>
		<tr>
			<th>Name</th>
			<th>Description</th>
			<th>Website</th>
			<th>GPS</th>
			<th>Actions</th>
		</tr>
	</thead>

	<tbody>
		{{range $place := .places}}
		<tr>
			<td>{{$place.Name}}</td>
			<td>{{$place.Description}}</td>
			<td>{{$place.Website}}</td>
			<td>{{$place.Lat}}, {{$place.Long}}</td>
			<td>
				<button class="pure-button">Edit</button>
				<button class="pure-button">Delete</button>
			</td>
		</tr>
		{{end}}
	</tbody>
</table>