import axios from 'axios';

interface Pet {
    ascii: string,
    description: string;
}

export async function fetchPet(): Promise<Pet> {
    try {
        const { data } = await axios.get('/v1/pet');

        if (!data || typeof data !== 'object') {
            return { ascii: '', description: '' };
        }

        return {
            ascii: data.ascii ?? '',
            description: data.description ?? '',
        };
    } catch (error: any) {
        if (axios.isAxiosError(error) && error.response?.status === 404) {
            return { ascii: '', description: '' };
        }
        throw error;
    }
}


export async function updatePet(pet: Pet): Promise<void> {
    await axios.put('/v1/pet', pet);    
}