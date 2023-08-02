export async function load({ fetch, params }) {
    const res = await fetch(`http://localhost:5000/api/systems`);
    const systems = await res.json();

    return { systems };
}