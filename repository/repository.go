/*
Dear Programmers,

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*                                                 *
*	This file belongs to Kevin Veros Hamonangan   *
*	and	Fandi Fladimir Dachi and is a part of     *
*	our	last project as the student of Del        *
*	Institute of Technology, Sitoluama.           *
*	Please contact us via Instagram:              *
*	sleepingnext and fandi_dachi                  *
*	before copying this file.                     *
*	Thank you, buddy. 😊                          *
*                                                 *
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*/

package repository

import (
	authPB "github.com/ta04/auth-service/model/proto"
)

// Repository is the interface of repositories.
// As there are number of repositories can be used.
type Repository interface {
	GetOneByUsername(*authPB.Auth2) (*authPB.Auth1, error)
	CreateOne(*authPB.Auth1) error
}
