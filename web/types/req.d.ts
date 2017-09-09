declare namespace Req {
    interface TabModel {
        id: string
        isShow: boolean
        label: string
    }

    interface ListModel {
        isDisable: boolean
        id: string
        key: string
        value: string
        desc: string
    }

    interface RequestModel {
        method: string
        url: string
        headers: ParamModel
        body: string
    }

    interface ResponseModel {
        status: string
        statusCode: number
        duration: string
        proto: string
        headers: ParamModel
        cookies: CookieModel[]
        body: string
        code: string
    }

    interface ParamModel {
        [key: string]: string[]
    }

    interface HeaderModel {
        key: string
        value: string
    }

    interface CookieModel {
        Name: string
        Value: string

        Path: string    // optional
        Domain: string    // optional
        Expires: string // optional
        RawExpires: string    // for reading cookies only

        // MaxAge=0 means no 'Max-Age' attribute specified.
        // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
        // MaxAge>0 means Max-Age attribute present and given in seconds
        MaxAge: number
        Secure: boolean
        HttpOnly: boolean
        Raw: string
        Unparsed: string[] // Raw text of unparsed attribute-value pairs
    }
}


        // isSending: boolean


        // reqMode: string
        // contentType: string
        // headers: ListModel[]
        // params: ListModel[]
        // bodys: ListModel[]
        // body: string
        // bodyMode: 'normal' | 'raw'

        // req: RequestModel
        // res: ResponseModel | undefined