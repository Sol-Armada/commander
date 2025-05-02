export async function authenticate(code) {
    const url = `${import.meta.env.VITE_API_BASE_URL}/login?code=${code}`;
    const headers = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
    };
    const res = await fetch(url, {
        method: 'GET',
        headers: headers,
    });
    const resJson = await res.json();

    if (res.status !== 200) {
        throw new Error(resJson.message || 'Authentication failed');
    }

    return resJson.data;
}

export async function authenticated(jwtToken) {
    if (!jwtToken) {
        return false;
    }

    const url = `${import.meta.env.VITE_API_BASE_URL}/auth`;
    const headers = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': `Bearer ${jwtToken}`,
    };
    const res = await fetch(url, {
        method: 'GET',
        headers: headers,
    });
    if (res.status !== 200) {
        return false;
    }
    return true;
}

export async function createOperation() {
    if (!localStorage.getItem('authenticated')) {
        throw new Error('User is not authenticated');
    }

    const url = `${import.meta.env.VITE_API_BASE_URL}/operations`;
    const headers = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('token')}`,
    };
    const res = await fetch(url, {
        method: 'POST',
        headers: headers,
    });
    if (res.status !== 200) {
        throw new Error('Failed to create operation');
    }
    const data = await res.json();
    return data.data;
}

export async function getOperations() {
    if (!localStorage.getItem('authenticated')) {
        throw new Error('User is not authenticated');
    }

    const url = `${import.meta.env.VITE_API_BASE_URL}/operations`;
    const headers = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('token')}`,
    };
    const res = await fetch(url, {
        method: 'GET',
        headers: headers,
    });
    if (res.status !== 200) {
        throw new Error('Failed to fetch operations data');
    }
    const data = await res.json();
    return data.data;
}

export async function getOperation(id) {
    if (!localStorage.getItem('authenticated')) {
        throw new Error('User is not authenticated');
    }

    const url = `${import.meta.env.VITE_API_BASE_URL}/operations/${id}`;
    const headers = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('token')}`,
    };
    const res = await fetch(url, {
        method: 'GET',
        headers: headers,
    });
    if (res.status !== 200) {
        throw new Error('Failed to fetch operation data');
    }
    const data = await res.json();
    return data.data;
}

export async function updateOperation(id, operation) {
    if (!localStorage.getItem('authenticated')) {
        throw new Error('User is not authenticated');
    }

    const url = `${import.meta.env.VITE_API_BASE_URL}/operations/${id}`;
    const headers = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('token')}`,
    };
    const res = await fetch(url, {
        method: 'PUT',
        headers: headers,
        body: JSON.stringify(operation),
    });
    if (res.status !== 200) {
        throw new Error('Failed to update operation data');
    }
    const data = await res.json();
    return data.data;
}

export async function getMember(id) {
    if (!localStorage.getItem('authenticated')) {
        throw new Error('User is not authenticated');
    }

    const url = `${import.meta.env.VITE_API_BASE_URL}/members/${id}`;
    const headers = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('token')}`,
    };
    const res = await fetch(url, {
        method: 'GET',
        headers: headers,
    });
    if (res.status !== 200) {
        throw new Error('Failed to fetch member data');
    }
    const data = await res.json();
    return data.data;
}

export async function getMembers() {
    if (!localStorage.getItem('authenticated')) {
        throw new Error('User is not authenticated');
    }

    const url = `${import.meta.env.VITE_API_BASE_URL}/members`;
    const headers = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('token')}`,
    };
    const res = await fetch(url, {
        method: 'GET',
        headers: headers,
    });
    if (res.status !== 200) {
        throw new Error('Failed to fetch members data');
    }
    const data = await res.json();
    return data.data;
}

export async function getShips() {
    if (!localStorage.getItem('authenticated')) {
        throw new Error('User is not authenticated');
    }

    const url = `${import.meta.env.VITE_API_BASE_URL}/ships`;
    const headers = {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': `Bearer ${localStorage.getItem('token')}`,
    };
    const res = await fetch(url, {
        method: 'GET',
        headers: headers,
    });
    if (res.status !== 200) {
        throw new Error('Failed to fetch ships data');
    }
    const data = await res.json();
    return data.data;
}
