import { supabase, api } from '../services/api'

export const getPosts = async () => {   
    //const t1 = await supabase.from("posts").insert({"title":"t1", "description": "d1", "body": "b1"})
    const {data} = await supabase.from("posts").select(); // .get('/posts'); 
    
    //const {data} = await api.get("/posts");
    //console.log("xix", xxx);

    if(data){
        return data;
    }

    return []
}

export const getPostBySlug = async (id) => {

    //TODO: BUSCAR UM POST EM ESPECIFICO.
    //const {data} = await api.get(`/posts?id=eq.${id}`)
    const {error, data} = await supabase.from("posts").select().eq("id", id);
    //console.log('get', data, error)
    if (data.length > 0) {
        return data[0]
    }
    return {} //"id":"","title":"","descritpion":"","body":"","created_at":""}
}