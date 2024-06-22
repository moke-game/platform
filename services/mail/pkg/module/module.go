package module

import (
	"go.uber.org/fx"

	"github.com/gstones/platform/services/mail/internal/service/private"
	"github.com/gstones/platform/services/mail/internal/service/public"
	"github.com/gstones/platform/services/mail/pkg/mailfx"
)

var MailModule = fx.Module("mail",
	public.ServiceModule,
	private.Module,
	mailfx.MailSettingsModule,
)

var MailPrivateModule = fx.Module("mail_private",
	private.Module,
	mailfx.MailSettingsModule,
)

var MailClientModule = fx.Module("mail_client",
	mailfx.MailSettingsModule,
	mailfx.MailClientModule,
)

var MailAllClientModule = fx.Module("mail_all_client",
	mailfx.MailSettingsModule,
	mailfx.MailClientModule,
	mailfx.MailClientPrivateModule,
)

var MailAllModule = fx.Module("mail_all",
	public.ServiceModule,
	private.Module,
	mailfx.MailSettingsModule,
	mailfx.MailClientModule,
	mailfx.MailClientPrivateModule,
)

var MailClientPrivateModule = fx.Module("mail_client_private",
	mailfx.MailSettingsModule,
	mailfx.MailClientPrivateModule,
)
