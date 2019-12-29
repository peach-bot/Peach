import wtforms
import wtforms_dynamic_fields as wdf

class SettingsForm(wtforms.Form):
    """The plugin settings form."""

def createsettings(settings: dict):
    dynamic = wdf.WTFormsDynamicFields()
    for plugin in settings["plugins"]:
        for setting in settings["plugins"][plugin]:
            if settings["plugins"][plugin][setting]["type"] == "bool":
                setattr(SettingsForm, plugin+"."+setting.lower().replace(" ", "_", -1), wtforms.BooleanField(setting))
    settingsform = SettingsForm()
    return settingsform