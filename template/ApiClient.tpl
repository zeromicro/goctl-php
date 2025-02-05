<?php

namespace {{.Namespace}};

class {{.ClientName}}Client extends ApiBaseClient
{
    {{range $i, $r := .Routes}}
    public function {{$r.HttpMethod}}{{$r.ActionPrefix}}{{$r.ActionName}}(
        {{if $r.RequestType}}$request, {{end}}
        $body=null
    ) {
        $result = $this->request(
            '{{$r.Prefix}}{{$r.UrlPath}}',
            '{{$r.HttpMethod}}',
            {{if $r.RequestHasPathParams}}$request->getPath(){{else}}null{{end}},
            {{if $r.RequestHasQueryString}}$request->getQuery(){{else}}null{{end}},
            {{if $r.RequestHasHeaders}}$request->getHeader(){{else}}null{{end}},
            {{if $r.RequestHasBody}}$body ?? $request->getBody(){{else}}$body{{end}}
        );
        {{if $r.ResponseType}}
        $response = new {{$r.ResponseType}}();
        {{if $r.ResponseHeaders}}
        $response->getHeader()
            {{range $n, $k := $r.ResponseHeaders}}->set{{$n}}($result['body']['{{$k}}'])
            {{end}};
        {{end}}
        {{if $r.ResponseBody}}
        $response->getBody()
            {{range $n, $k := $r.ResponseBody}}->set{{$n}}($result['body']['{{$k}}'])
            {{end}};
        {{end}}
        return $response;
        {{else}}
        return $result;
        {{end}}
    }
    {{end}}
}