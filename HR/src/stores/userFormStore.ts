import { defineStore } from 'pinia';

export const useUserFormStore = defineStore('userForm', {
    state: () => ({
        formData: {
            specialty: '',
            phone: '',
            experience: '',
            add_experience: '',
            links: '',
            filesList: [] as File[],
            agreed: false,
            isClosed: false,
        },
    }),
    actions: {
        setFormData(data: Partial<typeof this.formData>) {
            this.formData = { ...this.formData, ...data };
        },
    },
});