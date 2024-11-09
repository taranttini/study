const supabaseUrl = null;
const supabaseKey = null;

if (!supabaseUrl || !supabaseKey) {
    throw new Error("NECESSARIO INFORMAR A URL DO SUPABASE E A CHAVE, criar banco e token em 'https://supabase.com/'");
}

import { createClient } from '@supabase/supabase-js'
export const supabase = createClient(supabaseUrl, supabaseKey)

import axios from 'axios';
export const api = axios.create({
    baseURL: supabaseUrl + "/rest/v1/",
    headers: {
        apikey: supabaseKey,
        authorization: "Bearer " + supabaseKey,
        contentType: "application/json"
    }
})
