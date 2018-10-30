package solution

import "strings"

func numUniqueEmails(emails []string) int {
	cnt := make(map[string]bool)
	for _, email := range emails {
		slc := strings.Split(email, "@")
		local, domain := slc[0], slc[1]
		local = local[:strings.IndexRune(local, '+')]
		if local == "" {
			continue
		}
		local = strings.Replace(local, ".", "", -1)
		cnt[local+domain] = true
	}
	return len(cnt)
}
