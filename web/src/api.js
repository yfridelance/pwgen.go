import axios from 'axios';

const api = axios.create({
    baseURL: '/api',
    headers: {
        'Content-Type': 'application/json',
    },
});

export const generatePasswords = async (config) => {
    try {
        const response = await api.post('/generate', config);
        return response.data;
    } catch (error) {
        if (error.response) {
            throw new Error(error.response.data.error || 'Generation failed');
        }
        throw error;
    }
};
