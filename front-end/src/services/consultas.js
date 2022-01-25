import { api } from './config'

export default {
    listarPorPaciente:() => {
        return api.get('consultas/client/1')
    },

    listarPorMedico:(medicoId) => {
        return api.get(`consultas/medic/${medicoId}`)
    }
}