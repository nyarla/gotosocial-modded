// GoToSocial
// Copyright (C) GoToSocial Authors admin@gotosocial.org
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package email

const (
	testTemplate = "email_test.tmpl"
	testSubject  = "GoToSocial Test Email"
)

type TestData struct {
	// Username of admin user who sent the test.
	SendingUsername string
	// (Optional) message to include in the email.
	Message string
	// URL of the instance to present to the receiver.
	InstanceURL string
	// Name of the instance to present to the receiver.
	InstanceName string
}

func (s *sender) SendTestEmail(toAddress string, data TestData) error {
	return s.sendTemplate(testTemplate, testSubject, data, toAddress)
}
