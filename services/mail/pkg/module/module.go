package module

import (
	"go.uber.org/fx"

	"github.com/moke-game/platform/services/mail/internal/service/private"
	"github.com/moke-game/platform/services/mail/internal/service/public"
	"github.com/moke-game/platform/services/mail/pkg/mailfx"
)

var settings = mailfx.MailSettingsModule

var MailModule = fx.Module("mail",
	settings,
	public.ServiceModule,
	private.Module,
)

var MailPrivateModule = fx.Module("mail_private",
	settings,
	private.Module,
)

var MailClientModule = fx.Module("mail_client",
	settings,
	mailfx.MailClientModule,
)

var MailClientPrivateModule = fx.Module("mail_client_private",
	settings,
	mailfx.MailClientPrivateModule,
)

var MailAllClientModule = fx.Module("mail_all_client",
	settings,
	mailfx.MailClientModule,
	mailfx.MailClientPrivateModule,
)

var MailAllModule = fx.Module("mail_all",
	settings,
	public.ServiceModule,
	private.Module,
	mailfx.MailClientModule,
	mailfx.MailClientPrivateModule,
)
