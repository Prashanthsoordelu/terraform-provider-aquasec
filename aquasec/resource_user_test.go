package aquasec

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAquasecUserManagement(t *testing.T) {
	userID := "terraform-acctest"
	password := "password"
	name := "terraform acc testuser"
	email := "terraform@acctest.com"
	role := "administrator"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAquasecUser(userID, password, name, email, role),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAquasecUsersExists("resource.aquasec_user.new"),
				),
			},
		},
	})
}

func testAccCheckAquasecUser(userID string, password string, name string, email string, role string) string {
	return fmt.Sprintf(`
	resource "aquasec_user" "new" {
		user_id  = "%s"
		password = "%s"
		name     = "%s"
		email    = "%s"
		roles = [
		  "%s"
		]
	  }`, userID, password, name, email, role)

}

func testAccCheckAquasecUsersExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return NewNotFoundErrorf("%s in state", n)
		}

		if rs.Primary.ID == "" {
			return NewNotFoundErrorf("ID for %s in state", n)
		}

		return nil
	}
}
