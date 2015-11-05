/*
!!! DO NOT MODIFY !!!

autogenerated
 src: gen_list_job.tmpl
 job: &main.GenListJob{Name:"Balance"}
  on: Thu Nov 05 15:22:29 +0700 2015
  by: chakrit
*/
package omise

// BalanceList represents the list structure returned by Omise's REST API that contains
// Balance struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type BalanceList struct {
	List
	Data []*Balance `json:"data"`
}

// Find finds and returns Balance with the given id. Returns nil if not found.
func (list *BalanceList) Find(id string) *Balance {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}