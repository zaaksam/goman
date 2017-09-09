import _ from 'lodash'
import Axios, { AxiosRequestConfig } from 'axios'

interface APIResult<T> {
    code: number
    msg: string
    data?: T
}

interface APIStatic {
    /**
     * get请求
     */
    get<T>(url: string): Promise<APIResult<T>>
    /**
     * post请求
     */
    post<T>(url: string, data?: URLSearchParams): Promise<APIResult<T>>
}

async function request<T>(method: 'get' | 'post', url: string, data?: URLSearchParams) {
    let result: APIResult<T> = { code: -1, msg: '' }

    let conf: AxiosRequestConfig = {
        method: method,
        url: '/api' + url
        // paramsSerializer: (params) => {
        //     let str = ''
        //     if (_.size(params) > 0) {
        //         let i = 0
        //         _.forEach(params, (v, k) => {
        //             if (i > 0) {
        //                 str += '&'
        //             }
        //             str += k + '=' + encodeURIComponent(v)
        //         })
        //     }
        //     return str
        // }
    }
    if (data) {
        conf.data = data
    }

    try {
        let res = await Axios.request(conf)
        result.code = res.data.code
        result.msg = res.data.msg
        result.data = res.data.data

    } catch (err) {
        result.msg = err.message
    }

    return result
}

const API: APIStatic = {
    get: <T>(url: string) => {
        return request<T>('get', url)
    },
    post: <T>(url: string, data?: URLSearchParams) => {
        return request<T>('post', url, data)
    }
}

export default API