import useSWR from 'swr';

const fetcher = (url: string) => fetch(url).then((res) => res.json());

function Pet() {
    // Change the URL to your API endpoint
    const { data, error } = useSWR('http://localhost:8080/pets/2', fetcher);

    if (error) return <div>Failed to load</div>;
    if (!data) return <div>Loading...</div>;

    // Render your pet data
    return (
        <div>
            <h1>{data.name}</h1>
        <p>Tag: {data.tag}</p>
    </div>
);
}

export default Pet;
