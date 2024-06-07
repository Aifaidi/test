
`))

func renderToken(w http.ResponseWriter, redirectURL, idToken, accessToken, refreshToken, claims string) {

        var claimsJson ClaimsJson // declaration claims en json
        claimsByte := []byte(claims)
        err := json.Unmarshal(claimsByte, &claimsJson) // parsin de claims
        if err != nil {
                log.Printf("Error while parsing json: %s", err)
        }

        renderTemplate(w, tokenTmpl, tokenTmplData{
                IDToken:      idToken,
                AccessToken:  accessToken,
                RefreshToken: refreshToken,
                RedirectURL:  redirectURL,
                Claims:       claims,
                Email:        claimsJson.Email,
        })
}

func renderTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
        err := tmpl.Execute(w, data)
        if err == nil {
                return
        }

        switch err := err.(type) {
        case *template.Error:
                // An ExecError guarantees that Execute has not written to the underlying reader.
                log.Printf("Error rendering template %s: %s", tmpl.Name(), err)

                // TODO(ericchiang): replace with better internal server error.
                http.Error(w, "Internal server error", http.StatusInternalServerError)
        default:
                // An error with the underlying write, such as the connection being
                // dropped. Ignore for now.
        }
}
