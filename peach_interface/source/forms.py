import flask_wtf, wtforms

class SettingsForm(flask_wtf.FlaskForm):
    """The plugin settings form."""
    submit = wtforms.SubmitField('Save Changes')

def createsettings(settings: dict, form):
    for plugin in settings["plugins"]:
        for setting in settings["plugins"][plugin]:
            if settings["plugins"][plugin][setting]["type"] == "bool":
                setattr(SettingsForm, plugin+"_"+setting.lower().replace(" ", "_", -1), wtforms.BooleanField(setting))
    settingsform = SettingsForm(form)
    return settingsform