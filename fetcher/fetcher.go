package mantela_fetcher

import (
	"encoding/json"
)

func ParseMantela(data []byte) (*Mantela, error) {
	var mantela Mantela
	if err := json.Unmarshal(data, &mantela); err != nil {
		return nil, err
	}
	return &mantela, nil
}

func MergeMantela(m1, m2 *Mantela) *Mantela {
	// This function should implement the logic to merge two Mantela structures.
	// If both the source and the diff contain the same information,
	// priority is given to the information in the destination (it is always overwritten by the diff).
	// If you want a more complex merge or recursive merge, you may want to add a library to take care of the merging.
	if m2.Version != nil {
		m1.Version = m2.Version
	}
	if m2.AboutMe != nil {
		if m1.AboutMe == nil {
			m1.AboutMe = &AboutMe{}
		}
		if m2.AboutMe.Name != nil {
			m1.AboutMe.Name = m2.AboutMe.Name
		}
		if m2.AboutMe.PreferredPrefix != nil {
			m1.AboutMe.PreferredPrefix = m2.AboutMe.PreferredPrefix
		}
		if m2.AboutMe.Identifier != nil {
			m1.AboutMe.Identifier = m2.AboutMe.Identifier
		}
		if m2.AboutMe.SIPUsername != nil {
			m1.AboutMe.SIPUsername = m2.AboutMe.SIPUsername
		}
		if m2.AboutMe.SIPPassword != nil {
			m1.AboutMe.SIPPassword = m2.AboutMe.SIPPassword
		}
		if m2.AboutMe.SIPServer != nil {
			m1.AboutMe.SIPServer = m2.AboutMe.SIPServer
		}
		if m2.AboutMe.SIPPort != nil {
			m1.AboutMe.SIPPort = m2.AboutMe.SIPPort
		}
		if m2.AboutMe.Unavailable != nil {
			m1.AboutMe.Unavailable = m2.AboutMe.Unavailable
		}
	}
	if m2.Extensions != nil {
		for _, ext := range m2.Extensions {
			found := false
			for i, existingExt := range m1.Extensions {
				if existingExt.Identifier != nil && ext.Identifier != nil && *existingExt.Identifier == *ext.Identifier {
					// Update existing extension
					if ext.Name != nil {
						m1.Extensions[i].Name = ext.Name
					}
					if ext.Extension != nil {
						m1.Extensions[i].Extension = ext.Extension
					}
					if ext.Type != nil {
						m1.Extensions[i].Type = ext.Type
					}
					if ext.TransferTo != nil {
						m1.Extensions[i].TransferTo = ext.TransferTo
					}
					found = true
					break
				}
			}
			if !found {
				// Add new extension
				m1.Extensions = append(m1.Extensions, ext)
			}
		}
	}
	if m2.Providers != nil {
		for _, prov := range m2.Providers {
			found := false
			for i, existingProv := range m1.Providers {
				if existingProv.Identifier != nil && prov.Identifier != nil && *existingProv.Identifier == *prov.Identifier {
					// Update existing provider
					if prov.Name != nil {
						m1.Providers[i].Name = prov.Name
					}
					if prov.Prefix != nil {
						m1.Providers[i].Prefix = prov.Prefix
					}
					if prov.Identifier != nil {
						m1.Providers[i].Identifier = prov.Identifier
					}
					if prov.Mantela != nil {
						m1.Providers[i].Mantela = prov.Mantela
					}
					found = true
					break
				}
			}
			if !found {
				// Add new provider
				m1.Providers = append(m1.Providers, prov)
			}
		}
	}
	return m1
}
